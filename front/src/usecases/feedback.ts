import { Feedback } from "~/domain/feedback";
import { Sentry } from "~/infrastructure/sentry";

export interface IFeedbackUseCase {
  sendFeedback(userId: string, feedback: Feedback): Promise<void>;
}

export class FeedbackUseCase implements IFeedbackUseCase {
  async sendFeedback(userId: string, feedback: Feedback): Promise<void> {
    await Sentry.sendFeedback(userId, feedback);
  }
}
