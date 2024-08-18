import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { App, ProtectRoute } from './App.tsx'
import './index.css'

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

// route management in react frontend
const router = createBrowserRouter([
  {
    path:'/',
    element: <h1>404 Not found</h1>,
    children: [
      {
        path:"home",
        element: <ProtectRoute component={<App />}/>
      }
    ]
  }
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
