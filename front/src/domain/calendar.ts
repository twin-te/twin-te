/**
 * iCal 購読でタグの付いたコースをどう出力するか。
 * - sync: そのまま同期する（対象タグは空）
 * - exclude: 対象タグを持つコースをカレンダーに入れない
 * - transparent: 対象タグを持つコースを「予定なし」（TRANSP:TRANSPARENT）にする
 */
export type IcalSubscriptionMode = "sync" | "exclude" | "transparent";

export type IcalSubscription = {
  url: string;
  mode: IcalSubscriptionMode;
  targetTagIds: string[];
};
