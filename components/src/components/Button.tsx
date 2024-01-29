import { JSX } from "solid-js"
import { A } from "@solidjs/router"
import "./Button.css"

type Props = {
  children?: JSX.Element
  outline?: boolean
}

type ButtonProps = Props & JSX.ButtonHTMLAttributes<HTMLButtonElement>

export const Button = ({
  children,
  outline = false,
  ...props
}: ButtonProps) => {
  return (
    <button {...props} class={`${outline && "outline"} button ${props.class}`}>
      {children}
    </button>
  )
}

type ExtButtonProps = { path: string } & Props &
  JSX.LinkHTMLAttributes<HTMLAnchorElement>

export const ExtButton = ({
  children,
  outline = false,
  path,
  ...props
}: ExtButtonProps) => {
  return (
    <A
      href={path}
      {...props}
      class={`${outline && "outline"} button ${props.class} flex justify-center items-center`}
    >
      {children}
    </A>
  )
}
