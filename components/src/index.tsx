/* @refresh reload */
import { render } from "solid-js/web"
import { Route, Router } from "@solidjs/router"
import { notFound } from "./pages/notFound"
import { home } from "./pages/home"
import { mypage } from "./pages/mypage"
import "./index.css"

const root = document.getElementById("root")

render(
  () => (
    <Router>
      <Route path="/" component={home} />
      <Route path="/mypage" component={mypage} />
      <Route path="/*" component={notFound} />
    </Router>
  ),
  root!,
)
