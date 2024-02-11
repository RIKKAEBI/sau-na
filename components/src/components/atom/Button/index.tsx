import { JSX } from "solid-js"
import "./index.css"

type Props = {
  children?: JSX.Element
  outline?: boolean
} & JSX.ButtonHTMLAttributes<HTMLButtonElement>

export const Button = ({ children, outline = false, ...props }: Props) => {
  return (
    <button {...props} class={`${outline && "outline"} button ${props.class}`}>
      {children}
    </button>
  )
}
