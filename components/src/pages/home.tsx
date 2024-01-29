import { LoginButton } from "../components/LoginButton"
import logo from "../../public/Kasu-chan.png"

export const home = () => {
  return (
    <>
      <img src={logo} alt="kasu Chan" class="w-[60%] min-w-80 mx-auto mt-12" />
      <h1 class="text-center text-7xl mt-12 text-blue-600 google-font">
        さう〜な
      </h1>
      <div class="mt-8">
        <p class="text-center mx-4 break-words">
          <span class="inline-block">サウナの習慣化と健康を紐づけて</span>
          <span class="inline-block">管理できるウェブアプリです。</span>
        </p>
      </div>
      <LoginButton class="mx-auto max-w-48 mt-8" />
      <footer class="mt-auto text-center">
        <a
          href="https://github.com/sau-na/sau-na"
          class="text-xs text-gray-700"
          target="_blank"
        >
          ©️ Kasu Chan Company
        </a>
      </footer>
    </>
  )
}
