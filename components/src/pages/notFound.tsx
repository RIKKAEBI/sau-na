import { Layout } from "../layout"
import { ExtButton } from "../components/Button"

export const notFound = () => {
  return (
    <Layout>
      <img
        src="/Kasu-chan.png"
        alt="kasu Chan"
        class="w-[40%] min-w-80 mx-auto mt-12"
      />
      <h1 class="text-center mt-8">
        <span>そんなページねーよ！！</span>
        <span>こっちがぴえんだわ！！</span>
      </h1>
      <ExtButton path="/" outline class="mt-8 mx-auto block">
        トップへ戻る
      </ExtButton>
    </Layout>
  )
}
