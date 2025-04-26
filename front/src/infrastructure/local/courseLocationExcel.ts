import * as XLSX from "xlsx";
import { CourseLocationInfo } from "~/domain/courseLocation";

export const getKdbClassroom = async (
  file: File
): Promise<CourseLocationInfo> => {
  const book = XLSX.read(await file.arrayBuffer());
  const sheet = book.Sheets[book.SheetNames[0]];

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const range = XLSX.utils.decode_range(sheet["!ref"]!);

  for (let row = range.s.r; row <= range.e.r; row++) {
    const cellAddress = { r: row, c: range.s.c };
    const cellRef = XLSX.utils.encode_cell(cellAddress);
    const cell = sheet[cellRef];

    if (cell?.v === "科目番号") {
      range.s.r = row;
      break;
    }
  }

  const records: { courseId: string; classroom: string }[] = XLSX.utils
    .sheet_to_json<{
      [key: string]: string;
    }>(sheet, { range })
    .filter((it) => it["教室"].trim() !== "" && it["科目番号"].trim() !== "")
    .map((it) => ({ courseId: it["科目番号"], classroom: it["教室"] }));

  const courseIdToClassroom = records.reduce((acc, item) => {
    acc[item.courseId] = item.classroom;
    return acc;
  }, {} as { [key: string]: string });

  return {
    uploadAt: new Date(),
    courseLocations: courseIdToClassroom,
  };
};
