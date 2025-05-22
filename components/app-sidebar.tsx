"use client"

import * as React from "react"
import {
  IconCamera,
  IconChartBar,
  IconDashboard,
  IconDatabase,
  IconFileAi,
  IconFileDescription,
  IconFileWord,
  IconFolder,
  IconHelp,
  IconInnerShadowTop,
  IconListDetails,
  IconReport,
  IconSearch,
  IconSettings,
  IconUsers,
  IconPlane,
  IconPlaneDeparture,
  IconTicket,
  IconList,
  IconChartPie,
  IconAdjustments,
} from "@tabler/icons-react"

import { NavDocuments } from "@/components/nav-documents"
import { NavMain } from "@/components/nav-main"
import { NavSecondary } from "@/components/nav-secondary"
import { NavUser } from "@/components/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"

interface UserData {
  id: number
  username: string
  email: string
  role: string
  created_at: string
  updated_at: string
}

const defaultUserData = {
  name: "Guest",
  email: "guest@example.com",
  avatar: "/avatars/default.jpg",
  role: "GUEST"
}

const defaultNavData = {
  user: defaultUserData,
  navMain: [
    {
      title: "Dashboard",
      url: "/dashboard",
      icon: IconDashboard,
    },
    {
      title: "Danh sách chuyến bay",
      url: "/list",
      icon: IconPlane,
    },
    {
      title: "Thêm chuyến bay",
      url: "/flight",
      icon: IconPlaneDeparture,
    },
    {
      title: "Bán vé",
      url: "/ticket",
      icon: IconTicket,
    },
    {
      title: "Danh sách vé",
      url: "/list-ticket",
      icon: IconList,
    },
    {
      title: "Báo cáo doanh thu",
      url: "/report",
      icon: IconChartPie,
    },
    {
      title: "Thay đổi quy định",
      url: "/regulation",
      icon: IconAdjustments,
    },
  ],
  navClouds: [
    {
      title: "Capture",
      icon: IconCamera,
      isActive: true,
      url: "#",
      items: [
        {
          title: "Active Proposals",
          url: "#",
        },
        {
          title: "Archived",
          url: "#",
        },
      ],
    },
    {
      title: "Proposal",
      icon: IconFileDescription,
      url: "#",
      items: [
        {
          title: "Active Proposals",
          url: "#",
        },
        {
          title: "Archived",
          url: "#",
        },
      ],
    },
    {
      title: "Prompts",
      icon: IconFileAi,
      url: "#",
      items: [
        {
          title: "Active Proposals",
          url: "#",
        },
        {
          title: "Archived",
          url: "#",
        },
      ],
    },
  ],
  navSecondary: [],
  documents: [
  ],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const [data, setData] = React.useState(defaultNavData)

  React.useEffect(() => {
    const getNavData = () => {
      const userStr = localStorage.getItem('user')
      let userData = defaultUserData

      if (userStr) {
        try {
          const user: UserData = JSON.parse(userStr)
          userData = {
            name: user.username,
            email: user.email,
            role: user.role,
            avatar: "/avatars/default.jpg"
          }
        } catch (error) {
          console.error('Error parsing user data:', error)
        }
      }

      return {
        ...defaultNavData,
        user: userData
      }
    }

    setData(getNavData())
  }, [])

  return (
    <Sidebar collapsible="offcanvas" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              asChild
              className="data-[slot=sidebar-menu-button]:!p-1.5"
            >
              <a href="#">
                <IconInnerShadowTop className="!size-5" />
                <span className="text-base font-semibold">Acme Inc.</span>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        <NavDocuments items={data.documents} />
        <NavSecondary items={data.navSecondary} className="mt-auto" />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
    </Sidebar>
  )
}
