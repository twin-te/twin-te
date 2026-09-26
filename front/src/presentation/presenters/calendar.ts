/**
 * Return the iCal subscription URL that exports only the courses with at least one of the given tags.
 * The calendar endpoint reads the tag IDs from the repeated `tags[]` query parameter.
 * @param url - The iCal subscription URL, which exports all the registered courses.
 * @param tagIds - The IDs of the tags to be exported.
 */
export const icalUrlWithTags = (url: string, tagIds: string[]): string => {
  const icalUrl = new URL(url);
  icalUrl.searchParams.delete("tags[]");
  tagIds.forEach((tagId) => icalUrl.searchParams.append("tags[]", tagId));
  return icalUrl.toString();
};
