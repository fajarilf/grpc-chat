import { BackgroundType, type RoomBackground } from "$lib/types/room";

export const presetColors: { label: string; value: string }[] = [
  { label: "Slate", value: "#475569" },
  { label: "Sky", value: "#0284c7" },
  { label: "Emerald", value: "#059669" },
  { label: "Amber", value: "#d97706" },
  { label: "Rose", value: "#e11d48" },
  { label: "Violet", value: "#7c3aed" },
  { label: "Ink", value: "#14110d" },
  { label: "Teal", value: "#0d9488" },
];

export const presetImages: { label: string; src: string }[] = [
  { label: "Mountains", src: "https://picsum.photos/seed/mountains/640/480" },
  { label: "City", src: "https://picsum.photos/seed/city/640/480" },
  { label: "Ocean", src: "https://picsum.photos/seed/ocean/640/480" },
  { label: "Forest", src: "https://picsum.photos/seed/forest/640/480" },
  { label: "Desert", src: "https://picsum.photos/seed/desert/640/480" },
  { label: "Aurora", src: "https://picsum.photos/seed/aurora/640/480" },
];

export const defaultBackground: RoomBackground = {
  kind: BackgroundType.COLOR,
  value: presetColors[0].value,
};
