export type ColorHex = `#${string}`;
export type PresetColor =
  | "default"
  | "rose"
  | "orange"
  | "yellow"
  | "green"
  | "blue"
  | "purple";
export type TagColor = PresetColor | ColorHex;

export type Tag = {
  id: string;
  name: string;
  order: number; // 0-indices
  color: TagColor | null;
};

export type UndefinedTagOrder = -1;

export const undefinedTagOrder: UndefinedTagOrder = -1;
