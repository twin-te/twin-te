declare module "vue" {
  import type { Router, RouteLocationNormalized } from "vue-router";

  interface Vue {
    $router: Router;
    $route: RouteLocationNormalized;
  }
}
