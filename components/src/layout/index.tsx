import { JSX } from "solid-js"
import { Header } from "./Header"
import { Footer } from "./Footer"

type Props = {
  children: JSX.Element
}

export const Layout = ({ children }: Props) => {
  return (
    <>
      <Header />
      <div>{children}</div>
      <Footer />
    </>
  )
}
