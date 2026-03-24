import { captureFeedback } from "@sentry/vue";
import { Feedback } from "~/domain/feedback";

export class Sentry {
  static async sendFeedback(userId: string, feedback: Feedback) {
    return captureFeedback(
      {
        email: feedback.email,
        name: userId,
        message: feedback.content,
        tags: {
          type: feedback.type,
        },
      },
      {
        attachments: await Promise.all(
          feedback.screenShots.map(async (s) => ({
            filename: s.name,
            data: await s.arrayBuffer().then((buf) => new Uint8Array(buf)),
          }))
        ),
        includeReplay: true,
      }
    );
  }
}
