import { JSX } from "solid-js"
import logo from "../assets/web_light_rd_SI@4x.png"
import { ExtButton } from "./Button"

export const LoginButton = (
  props: JSX.ButtonHTMLAttributes<HTMLAnchorElement>,
) => {
  return (
    <ExtButton path="/register" {...props}>
      <img src={logo} alt="google" />
    </ExtButton>
  )
}
