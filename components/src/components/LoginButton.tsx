import { JSX } from "solid-js"
import logo from "../assets/web_light_rd_SI@4x.png"

export const LoginButton = (
  props: JSX.ButtonHTMLAttributes<HTMLButtonElement>,
) => {
  return (
    <button {...props}>
      <img src={logo} alt="google" />
    </button>
  )
}
