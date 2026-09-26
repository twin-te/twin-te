import { icalUrlWithTags } from "../presenters/calendar";

const url =
  "https://app.twinte.net/calendar/v1/timetable.ics?token=0f6c6d6e-1d0a-4b8e-9a51-8f0d6c3c2b1a";
const tagA = "3b1f8a8e-0c9d-4a57-8f36-2f7c1e9d5a10";
const tagB = "c2d4e6f8-1a3b-4c5d-8e7f-9a0b1c2d3e4f";

describe(icalUrlWithTags.name, () => {
  it("appends each tag ID as a tags[] query parameter.", () => {
    const result = new URL(icalUrlWithTags(url, [tagA, tagB]));
    expect(result.searchParams.getAll("tags[]")).toEqual([tagA, tagB]);
  });

  it("keeps the token query parameter.", () => {
    const result = new URL(icalUrlWithTags(url, [tagA]));
    expect(result.searchParams.get("token")).toBe(
      "0f6c6d6e-1d0a-4b8e-9a51-8f0d6c3c2b1a"
    );
  });

  it("replaces the tags[] query parameters that the URL already has.", () => {
    const result = new URL(
      icalUrlWithTags(icalUrlWithTags(url, [tagA]), [tagB])
    );
    expect(result.searchParams.getAll("tags[]")).toEqual([tagB]);
  });

  it("encodes the brackets of tags[].", () => {
    expect(icalUrlWithTags(url, [tagA])).toBe(`${url}&tags%5B%5D=${tagA}`);
  });
});
