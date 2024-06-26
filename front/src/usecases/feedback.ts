import axios from "axios";
import {
  InternalServerError,
  NetworkError,
  UnauthenticatedError,
  isResultError,
} from "~/domain/error";
import { Feedback } from "~/domain/feedback";
import { Firebase } from "~/infrastructure/firebase";

export interface IFeedbackUseCase {
  sendFeedback(
    userId: string,
    feedback: Feedback
  ): Promise<null | UnauthenticatedError | NetworkError | InternalServerError>;
}

export class FeedbackUseCase implements IFeedbackUseCase {
  #firebase: Firebase;

  constructor() {
    this.#firebase = Firebase.getInstance();
  }

  async sendFeedback(
    userId: string,
    feedback: Feedback
  ): Promise<null | UnauthenticatedError | NetworkError | InternalServerError> {
    const result = await this.#firebase.saveScreenshots(
      feedback.screenShots,
      userId
    );

    if (isResultError(result)) return result;
    const screenshotUrls = result;

    const formData = new FormData();
    formData.append("entry.1670691903", userId);
    formData.append("entry.1912135146", feedback.type);
    formData.append("entry.240515057", feedback.content);
    formData.append("entry.1261834897", feedback.email);
    formData.append("entry.1457360447", screenshotUrls.join(", "));

    return axios
      .post(
        "https://docs.google.com/forms/u/0/d/e/1FAIpQLSecBhx7GHm870_BHEjm94NAKOoxJPzbTL-pxpvy6BqxQQh6ag/formResponse",
        formData
      )
      .then(() => null)
      .catch(() => new InternalServerError("Faied to save feedback."));
  }
}
