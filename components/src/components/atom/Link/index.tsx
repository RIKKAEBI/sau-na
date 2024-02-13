import { JSX } from "solid-js"
import { A } from "@solidjs/router"
import "./index.css"

type Props = {
  path: string
  children?: JSX.Element
  outline?: boolean
} & JSX.LinkHTMLAttributes<HTMLAnchorElement>

export const Link = ({ children, outline = false, path, ...props }: Props) => {
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
