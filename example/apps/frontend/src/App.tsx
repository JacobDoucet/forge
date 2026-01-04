import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { Box } from '@mui/material'
import Layout from './components/Layout'
import TasksPage from './pages/TasksPage'
import ProjectsPage from './pages/ProjectsPage'

function App() {
  return (
    <BrowserRouter>
      <Box sx={{ display: 'flex', minHeight: '100vh' }}>
        <Layout>
          <Routes>
            <Route path="/" element={<Navigate to="/tasks" replace />} />
            <Route path="/tasks" element={<TasksPage />} />
            <Route path="/projects" element={<ProjectsPage />} />
          </Routes>
        </Layout>
      </Box>
    </BrowserRouter>
  )
}

export default App
