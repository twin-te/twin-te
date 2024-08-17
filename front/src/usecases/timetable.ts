import {
  createPromiseClient,
  ConnectError,
  Code,
  PromiseClient,
  Transport,
} from "@connectrpc/connect";
import { Mutex } from "async-mutex";
import { RegisteredCourse, Course } from "~/domain/course";
import {
  InternalServerError,
  NetworkError,
  NotFoundError,
  UnauthenticatedError,
  isResultError,
} from "~/domain/error";
import { Module, modules } from "~/domain/module";
import { NormalSchedule, Schedule, isNormalSchedule } from "~/domain/schedule";
import { Tag } from "~/domain/tag";
import {
  NormalTimetable,
  initializeTimetable,
  normalSchedulesToNormalTimetable,
  timetableToSchedules,
} from "~/domain/timetable";
import {
  toPBAcademicYear,
  toPBUUID,
} from "~/infrastructure/api/converters/shared";
import {
  fromPBCourse,
  fromPBRegisteredCourse,
  fromPBTag,
  toPBCourseMethod,
  toPBCredit,
  toPBInstructors,
  toPBSchedules,
} from "~/infrastructure/api/converters/timetablev1";
import { assurePresence } from "~/infrastructure/api/converters/utils";
import * as SharedPB from "~/infrastructure/api/gen/shared/type_pb";
import { TimetableService } from "~/infrastructure/api/gen/timetable/v1/service_connect";
import * as TimetableV1PB from "~/infrastructure/api/gen/timetable/v1/type_pb";
import { handleError } from "~/infrastructure/api/utils";
import {
  addElementsInArray,
  deepCopy,
  deleteElementInArray,
  updateElementInArray,
} from "~/utils";

export interface ITimetableUseCase {
  getCoursesByCodes(data: {
    year: number;
    codes: string[];
  }): Promise<
    Course[] | UnauthenticatedError | NetworkError | InternalServerError
  >;

  searchCourses(conds: {
    year: number;
    keywords: string[];
    codePrefixes: { included: string[]; excluded: string[] };
    schedules: {
      fullyIncluded: Schedule[];
      partiallyOverlapped: Schedule[];
    };
    offset: number;
    limit: number;
  }): Promise<
    Course[] | UnauthenticatedError | NetworkError | InternalServerError
  >;

  searchCoursesOnBlank(conds: {
    year: number;
    keywords: string[];
    codePrefixes: { included: string[]; excluded: string[] };
    offset: number;
    limit: number;
  }): Promise<
    Course[] | UnauthenticatedError | NetworkError | InternalServerError
  >;

  addCoursesByCodes(data: {
    year: number;
    codes: string[];
  }): Promise<
    | RegisteredCourse[]
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  addCourseManually(
    data: Pick<
      RegisteredCourse,
      | "year"
      | "name"
      | "instructors"
      | "credit"
      | "schedules"
      | "methods"
      | "rooms"
    >
  ): Promise<
    RegisteredCourse | UnauthenticatedError | NetworkError | InternalServerError
  >;

  getRegisteredCourses(
    year?: number,
    tagID?: string
  ): Promise<
    | RegisteredCourse[]
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  getRegisteredCourseById(
    id: string
  ): Promise<
    | RegisteredCourse
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  updateRegisteredCourse(
    id: string,
    data: Partial<Omit<RegisteredCourse, "id" | "year" | "code">>
  ): Promise<
    | RegisteredCourse
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  deleteRegisteredCourse(
    id: string
  ): Promise<
    | null
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  /**
   * Return true if the schedules do not overlap comparing to the schedules of registered courses. Return false otherwise.
   */
  checkScheduleDuplicate(
    year: number,
    schedules: Schedule[]
  ): Promise<
    boolean | UnauthenticatedError | NetworkError | InternalServerError
  >;

  createTag(
    name: string
  ): Promise<Tag | UnauthenticatedError | NetworkError | InternalServerError>;

  getTags(): Promise<
    Tag[] | UnauthenticatedError | NetworkError | InternalServerError
  >;

  getTagById(
    id: string
  ): Promise<
    | Tag
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  updateTagName(
    id: string,
    name: string
  ): Promise<
    | Tag
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  /**
   * Update tag orders. All tag ids that the user has must be specified.
   * @param ids - List of tag ids. The index represents each tag order.
   */
  updateTagOrders(
    ids: string[]
  ): Promise<Tag[] | UnauthenticatedError | NetworkError | InternalServerError>;

  deleteTag(
    id: string
  ): Promise<
    | null
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;
}

export class TimetableUseCase implements ITimetableUseCase {
  #client: PromiseClient<typeof TimetableService>;

  #mutex: {
    registeredCourses: Mutex;
    tags: Mutex;
  };

  #registeredCourses?: RegisteredCourse[];
  #tags?: Tag[];

  constructor(transport: Transport) {
    this.#client = createPromiseClient(TimetableService, transport);
    this.#mutex = {
      registeredCourses: new Mutex(),
      tags: new Mutex(),
    };
  }

  async #getRegisteredCourses(): Promise<
    | RegisteredCourse[]
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#mutex.registeredCourses.runExclusive(() => {
      if (this.#registeredCourses) {
        return this.#registeredCourses;
      }

      return this.#client
        .getRegisteredCourses({})
        .then((res) => res.registeredCourses.map(fromPBRegisteredCourse))
        .then((registeredCourses) => {
          return (this.#registeredCourses = registeredCourses);
        })
        .catch((error) => {
          return handleError(error, (connectError: ConnectError) => {
            if (connectError.code === Code.Unauthenticated) {
              return new UnauthenticatedError();
            }

            throw error;
          });
        });
    });
  }

  async #getTags(): Promise<
    Tag[] | UnauthenticatedError | NetworkError | InternalServerError
  > {
    return this.#mutex.tags.runExclusive(() => {
      if (this.#tags) {
        return this.#tags;
      }

      return this.#client
        .getTags({})
        .then((res) => res.tags.map(fromPBTag))
        .then((tags) => {
          return (this.#tags = tags);
        })
        .catch((error) => {
          return handleError(error, (connectError: ConnectError) => {
            if (connectError.code === Code.Unauthenticated) {
              return new UnauthenticatedError();
            }

            throw error;
          });
        });
    });
  }

  async getCoursesByCodes(data: {
    year: number;
    codes: string[];
  }): Promise<
    Course[] | UnauthenticatedError | NetworkError | InternalServerError
  > {
    const pbAcademicYear = toPBAcademicYear(data.year);
    return this.#client
      .getCoursesByCodes({ year: pbAcademicYear, codes: data.codes })
      .then((res) => res.courses.map(fromPBCourse))
      .catch((error) => handleError(error));
  }

  async searchCourses(conds: {
    year: number;
    keywords: string[];
    codePrefixes: { included: string[]; excluded: string[] };
    schedules: {
      fullyIncluded: Schedule[];
      partiallyOverlapped: Schedule[];
    };
    offset: number;
    limit: number;
  }): Promise<
    Course[] | UnauthenticatedError | NetworkError | InternalServerError
  > {
    return this.#client
      .searchCourses({
        year: toPBAcademicYear(conds.year),
        keywords: conds.keywords,
        codePrefixesIncluded: conds.codePrefixes.included,
        codePrefixesExcluded: conds.codePrefixes.excluded,
        schedulesFullyIncluded: toPBSchedules(
          conds.schedules.fullyIncluded,
          []
        ),
        schedulesPartiallyOverlapped: toPBSchedules(
          conds.schedules.partiallyOverlapped,
          []
        ),
        offset: conds.offset,
        limit: conds.limit,
      })
      .then((res) => res.courses.map(fromPBCourse))
      .catch((error) => handleError(error));
  }

  async searchCoursesOnBlank(conds: {
    year: number;
    keywords: string[];
    codePrefixes: { included: string[]; excluded: string[] };
    offset: number;
    limit: number;
  }): Promise<
    Course[] | UnauthenticatedError | NetworkError | InternalServerError
  > {
    const result = await this.getRegisteredCourses(conds.year);
    if (isResultError(result)) {
      return result;
    }

    const timetable = initializeTimetable(modules, true);

    result
      .map(({ schedules }) => schedules)
      .flat()
      .filter(isNormalSchedule)
      .forEach(({ module, day, period }) => {
        timetable.normal[module][day][period] = false;
      });

    const schedules = timetableToSchedules(timetable);

    return this.#client
      .searchCourses({
        year: toPBAcademicYear(conds.year),
        keywords: conds.keywords,
        codePrefixesIncluded: conds.codePrefixes.included,
        codePrefixesExcluded: conds.codePrefixes.excluded,
        schedulesFullyIncluded: toPBSchedules(schedules, []),
        schedulesPartiallyOverlapped: [],
        offset: conds.offset,
        limit: conds.limit,
      })
      .then((res) => res.courses.map(fromPBCourse))
      .catch((error) => handleError(error));
  }

  async addCoursesByCodes(data: {
    year: number;
    codes: string[];
  }): Promise<
    | RegisteredCourse[]
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    const pbAcademicYear = toPBAcademicYear(data.year);
    return this.#client
      .createRegisteredCoursesByCodes({
        year: pbAcademicYear,
        codes: data.codes,
      })
      .then((res) => res.registeredCourses.map(fromPBRegisteredCourse))
      .then((registeredCourses) => {
        return this.#mutex.registeredCourses.runExclusive(() => {
          if (this.#registeredCourses) {
            addElementsInArray(
              this.#registeredCourses,
              ...deepCopy(registeredCourses)
            );
          }
          return registeredCourses;
        });
      })
      .catch((error) =>
        handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          if (connectError.code === Code.NotFound) {
            return new NotFoundError();
          }

          throw error;
        })
      );
  }

  async addCourseManually(
    data: Pick<
      RegisteredCourse,
      | "year"
      | "name"
      | "instructors"
      | "credit"
      | "schedules"
      | "methods"
      | "rooms"
    >
  ): Promise<
    RegisteredCourse | UnauthenticatedError | NetworkError | InternalServerError
  > {
    return this.#client
      .createRegisteredCourseManually({
        year: toPBAcademicYear(data.year),
        name: data.name,
        instructors: toPBInstructors(data.instructors),
        credit: toPBCredit(data.credit),
        schedules: toPBSchedules(data.schedules, data.rooms),
        methods: data.methods.map(toPBCourseMethod),
      })
      .then((res) =>
        fromPBRegisteredCourse(assurePresence(res.registeredCourse))
      )
      .then((registeredCourse) => {
        return this.#mutex.registeredCourses.runExclusive(() => {
          if (this.#registeredCourses) {
            addElementsInArray(
              this.#registeredCourses,
              deepCopy(registeredCourse)
            );
          }
          return registeredCourse;
        });
      })
      .catch((error) =>
        handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          throw error;
        })
      );
  }

  async getRegisteredCourses(
    year?: number,
    tagID?: string
  ): Promise<
    | RegisteredCourse[]
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    const result = await this.#getRegisteredCourses();

    if (isResultError(result)) {
      return result;
    }

    let registeredCourses = result;

    if (year) {
      registeredCourses = registeredCourses.filter(
        (registeredCourse) => registeredCourse.year === year
      );
    }

    if (tagID) {
      registeredCourses = registeredCourses.filter((registeredCourse) =>
        registeredCourse.tagIds.includes(tagID)
      );
    }

    return deepCopy(registeredCourses);
  }

  async getRegisteredCourseById(
    id: string
  ): Promise<
    | RegisteredCourse
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    const result = await this.#getRegisteredCourses();

    if (isResultError(result)) {
      return result;
    }

    const registeredCourse = result.find(
      (registeredCourse) => registeredCourse.id === id
    );

    return registeredCourse ?? new NotFoundError();
  }

  // If you want to update either schedules or rooms, please specify both.
  async updateRegisteredCourse(
    id: string,
    data: Partial<Omit<RegisteredCourse, "id" | "year" | "code">>
  ): Promise<
    | RegisteredCourse
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#client
      .updateRegisteredCourse({
        id: toPBUUID(id),
        name: data.name,
        instructors: data.instructors
          ? toPBInstructors(data.instructors)
          : undefined,
        credit: data.credit ? toPBCredit(data.credit) : undefined,
        methods: data.methods
          ? new TimetableV1PB.CourseMethodList({
              values: data.methods.map(toPBCourseMethod),
            })
          : undefined,
        schedules:
          data.schedules && data.rooms
            ? new TimetableV1PB.ScheduleList({
                values: toPBSchedules(data.schedules, data.rooms),
              })
            : undefined,
        memo: data.memo,
        attendance: data.attendance,
        absence: data.absence,
        late: data.late,
        tagIds: data.tagIds
          ? new SharedPB.UUIDList({ values: data.tagIds.map(toPBUUID) })
          : undefined,
      })
      .then((res) =>
        fromPBRegisteredCourse(assurePresence(res.registeredCourse))
      )
      .then((registeredCourse) => {
        return this.#mutex.registeredCourses.runExclusive(() => {
          if (this.#registeredCourses) {
            updateElementInArray(
              this.#registeredCourses,
              deepCopy(registeredCourse)
            );
          }
          return registeredCourse;
        });
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          if (connectError.code === Code.NotFound) {
            return new NotFoundError();
          }

          throw error;
        });
      });
  }

  async deleteRegisteredCourse(
    id: string
  ): Promise<
    | null
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#client
      .deleteRegisteredCourse({ id: toPBUUID(id) })
      .then(() => {
        return this.#mutex.registeredCourses.runExclusive(() => {
          if (this.#registeredCourses) {
            deleteElementInArray(this.#registeredCourses, id);
          }
          return null;
        });
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          if (connectError.code === Code.NotFound) {
            return new NotFoundError();
          }

          throw error;
        });
      });
  }

  async checkScheduleDuplicate(
    year: number,
    schedules: Schedule[]
  ): Promise<
    boolean | UnauthenticatedError | NetworkError | InternalServerError
  > {
    const result = await this.getRegisteredCourses(year);
    if (isResultError(result)) return result;

    const registeredNormalSchedules: NormalSchedule[] = result
      .map(({ schedules }) => schedules)
      .flat()
      .filter(isNormalSchedule);

    const regisreredNormalTimetable: NormalTimetable<
      Module,
      boolean
    > = normalSchedulesToNormalTimetable(registeredNormalSchedules);

    return !schedules
      .filter(isNormalSchedule)
      .some(
        ({ module, day, period }) =>
          regisreredNormalTimetable[module][day][period]
      );
  }

  async createTag(
    name: string
  ): Promise<Tag | UnauthenticatedError | NetworkError | InternalServerError> {
    return this.#client
      .createTag({ name })
      .then((res) => fromPBTag(assurePresence(res.tag)))
      .then((tag) => {
        return this.#mutex.tags.runExclusive(() => {
          if (this.#tags) {
            addElementsInArray(this.#tags, deepCopy(tag));
          }
          return tag;
        });
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          throw error;
        });
      });
  }

  async getTags(): Promise<
    Tag[] | UnauthenticatedError | NetworkError | InternalServerError
  > {
    const result = await this.#getTags();

    if (isResultError(result)) {
      return result;
    }

    return deepCopy(result);
  }

  async getTagById(
    id: string
  ): Promise<
    | Tag
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    const result = await this.#getTags();

    if (isResultError(result)) {
      return result;
    }

    const tag = result.find((tag) => tag.id === id);
    return tag ?? new NotFoundError();
  }

  async updateTagName(
    id: string,
    name: string
  ): Promise<
    | Tag
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#client
      .updateTag({ id: toPBUUID(id), name })
      .then((res) => fromPBTag(assurePresence(res.tag)))
      .then((tag) => {
        return this.#mutex.tags.runExclusive(() => {
          if (this.#tags) {
            updateElementInArray(this.#tags, deepCopy(tag));
          }
          return tag;
        });
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          if (connectError.code === Code.NotFound) {
            return new NotFoundError();
          }

          throw error;
        });
      });
  }

  /**
   * Update tag orders. All tag ids that the user has must be specified.
   * @param ids - List of tag ids. The index represents each tag order.
   */
  async updateTagOrders(
    ids: string[]
  ): Promise<
    Tag[] | UnauthenticatedError | NetworkError | InternalServerError
  > {
    return this.#client
      .rearrangeTags({ ids: ids.map(toPBUUID) })
      .then((res) => res.tags.map((tag) => fromPBTag(assurePresence(tag))))
      .then((tags) => {
        return this.#mutex.tags.runExclusive(() => {
          if (this.#tags) {
            this.#tags = deepCopy(tags);
          }
          return tags;
        });
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          throw error;
        });
      });
  }

  async deleteTag(
    id: string
  ): Promise<
    | null
    | NotFoundError
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#client
      .deleteTag({ id: toPBUUID(id) })
      .then(() => {
        return this.#mutex.tags.runExclusive(() => {
          if (this.#tags) {
            deleteElementInArray(this.#tags, id);
          }
          return null;
        });
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }

          if (connectError.code === Code.NotFound) {
            return new NotFoundError();
          }

          throw error;
        });
      });
  }
}
