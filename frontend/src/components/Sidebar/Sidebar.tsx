import React from 'react'

interface SidebarProps {
    sideOpen: boolean;
}

const Sidebar: React.FC<SidebarProps> = ({sideOpen}) => {
  return (
    <div className={`sidebar-body ${sideOpen?"show":"hidden"}`}>
        Sidebar
    </div>
  )
}

export default Sidebar