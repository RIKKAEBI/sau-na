/* @refresh reload */
import { render } from "solid-js/web"
import { Route, Router } from "@solidjs/router"
import { notFound } from "./pages/notFound"
import { home } from "./pages/home"
import "./index.css"

const root = document.getElementById("root")

render(
  () => (
    <Router>
      <Route path="/" component={home} />
      <Route path="/*" component={notFound} />
    </Router>
  ),
  root!,
)
