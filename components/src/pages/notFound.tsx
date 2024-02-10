import { LinkButton } from "../components/atom/Button"
import { Layout } from "../layout"

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
      <LinkButton path="/" outline class="mt-8 mx-auto block">
        トップへ戻る
      </LinkButton>
    </Layout>
  )
}
