import type { Meta, StoryObj } from "storybook-solidjs"
import { Link } from "."
import { Route, Router } from "@solidjs/router"

const meta = {
  title: "atom/Link",
  component: Link,
  args: {
    children: "クリックして！",
    outline: true,
  },
  decorators: [
    (Story) => (
      <Router>
        <Route component={() => <Story path="#" />} />
      </Router>
    ),
  ],
} satisfies Meta<typeof Link>

export default meta
type Story = StoryObj<typeof meta>

export const Primary: Story = {}
