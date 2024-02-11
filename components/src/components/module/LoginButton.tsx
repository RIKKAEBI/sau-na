import { JSX } from "solid-js"
import logo from "../../assets/web_light_rd_SI@4x.png"
import { Link } from "../atom/Link"

export const LoginButton = (
  props: JSX.ButtonHTMLAttributes<HTMLAnchorElement>,
) => {
  return (
    <Link path="/register" {...props}>
      <img src={logo} alt="google" />
    </Link>
  )
}
