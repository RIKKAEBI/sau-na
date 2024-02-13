import type { Meta, StoryObj } from "storybook-solidjs"
import { Button } from "."

const meta = {
  title: "atom/Button",
  component: Button,
  args: {
    children: "クリックして！",
    outline: true,
  },
} satisfies Meta<typeof Button>

export default meta
type Story = StoryObj<typeof meta>

export const Primary: Story = {}
