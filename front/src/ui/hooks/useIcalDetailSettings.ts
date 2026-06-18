import { ref } from "vue";
import { IcalSubscriptionMode } from "~/domain/calendar";
import {
  InternalServerError,
  isResultError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import { calendarUseCase, timetableUseCase } from "~/usecases";
import type { IcalDetailSettingsValue } from "~/ui/components/IcalDetailSettings.vue";

/**
 * iCal 購読の「タグごとの表示設定」の状態を読み込み・保存するための composable。
 * PC のモーダルと SP のサブページの両方から利用する。
 */
export const useIcalDetailSettings = () => {
  const tags = ref<{ id: string; name: string }[]>([]);
  const value = ref<IcalDetailSettingsValue>({
    mode: "sync",
    targetTagIds: [],
  });
  const loaded = ref(false);

  const load = async () => {
    const [settings, tagList] = await Promise.all([
      calendarUseCase.getIcalSubscriptionUrl(),
      timetableUseCase.listTags(),
    ]);

    if (!isResultError(tagList)) {
      tags.value = tagList
        .sort((a, b) => a.order - b.order)
        .map(({ id, name }) => ({ id, name }));
    }

    if (!isResultError(settings) && "url" in settings && settings.url) {
      value.value = {
        mode: settings.mode,
        targetTagIds: settings.targetTagIds,
      };
    }

    loaded.value = true;
  };

  const save = (): Promise<
    null | UnauthenticatedError | NetworkError | InternalServerError
  > => {
    // 不変条件: 対象タグが無い場合は SYNC（そのまま同期）として保存する。
    const mode: IcalSubscriptionMode =
      value.value.targetTagIds.length === 0 ? "sync" : value.value.mode;
    const targetTagIds = mode === "sync" ? [] : value.value.targetTagIds;
    return calendarUseCase.updateIcalSubscription(mode, targetTagIds);
  };

  return { tags, value, loaded, load, save };
};
