import { onMount } from "solid-js"

export const register = () => {
  onMount(() => {
    window.setTimeout(() => {
      window.location.reload()
    }, 1)
  })

  return <></>
}
