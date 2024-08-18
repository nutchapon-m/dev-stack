import './App.css'
import { useState } from 'react'
import { Box, Paper, Button } from '@mui/material'
import { useLocation, Navigate } from 'react-router-dom'
import Sidebar from './components/Sidebar/Sidebar'

interface ProtectRouteProps {
  component: JSX.Element
}

export const ProtectRoute: React.FC<ProtectRouteProps> = ({component}) => {
  const isAuth: boolean = !!sessionStorage.getItem("session_token")    
  return (isAuth? component: <Navigate to={"/"} replace/>)
}


function Layout() {
  const [over, setOver] = useState(false)
  const [sideOpen, setSideOpen] = useState(false)
  const location = useLocation();


  if (location.pathname === "/") {
    const isAuthen = !!sessionStorage.getItem("session_token");
    return  (isAuthen ? <Navigate to={"/home"} replace/> : <Navigate to={"/"} replace/>);
  } else {
    return (
      <Box className="layout">
        <Box>
          <Sidebar sideOpen={sideOpen}/>
        </Box>  
        <Box 
          className="content" 
          component={Paper} 
          elevation={over?6:2} 
          sx={{
            margin: 4
          }}
          onMouseOver={()=>{setOver(true)}}
          onMouseLeave={()=>{setOver(false)}}
        >
          <Button onClick={()=>{setSideOpen(true)}}>
            open
          </Button>
          Content
        </Box>
      </Box>
    )
  } 
}


export function App() {
  return (
    <div>
      <Layout/>
    </div>
  )
}
