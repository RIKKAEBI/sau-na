export const apiClient = async (path: string) => {
  const res = await fetch(new URL(path, import.meta.env.VITE_API_URL))
  console.log(await res.json())
}
