import dayjs, { Dayjs } from 'dayjs';
import * as SharedPB from '../../api/gen/shared/type_pb';

export const fromPBUUID = (pbUUID: SharedPB.UUID): string => {
	return pbUUID.value;
};

export const fromPBRFC3339DateTime = (pbDateTime: SharedPB.RFC3339DateTime): Dayjs => {
	return dayjs(pbDateTime.value);
};
