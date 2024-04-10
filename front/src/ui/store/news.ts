import { computed, reactive } from "vue";
import { isResultError } from "~/domain/error";
import { News } from "~/domain/news";
import { newsUseCase } from "~/usecases";
import { deepCopy, updateAllElementsInArray } from "~/utils";

const news = reactive<News[]>([]);

const setNews = async () => {
  const result = await newsUseCase.getNews();
  if (isResultError(result)) throw result;
  updateAllElementsInArray(news, result);
};

const readNews = async (ids: string[]) => {
  const result = await newsUseCase.readNews(ids);
  if (isResultError(result)) throw result;
  await setNews();
};

const useNews = () => {
  return {
    news: computed(() => deepCopy(news)),
    unreadNews: computed(() => deepCopy(news.filter(({ read }) => !read))),
    readNews,
    initializeNews: setNews,
  };
};

export default useNews;
