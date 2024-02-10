import { JSX } from "solid-js"
import logo from "../../assets/web_light_rd_SI@4x.png"
import { LinkButton } from "../atom/Button"

export const LoginButton = (
  props: JSX.ButtonHTMLAttributes<HTMLAnchorElement>,
) => {
  return (
    <LinkButton path="/register" {...props}>
      <img src={logo} alt="google" />
    </LinkButton>
  )
}
