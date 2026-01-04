import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import {
  Box,
  Button,
  Card,
  CardContent,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Fab,
  IconButton,
  Stack,
  TextField,
  Typography,
  CircularProgress,
  Alert,
} from '@mui/material'
import AddIcon from '@mui/icons-material/Add'
import DeleteIcon from '@mui/icons-material/Delete'
import EditIcon from '@mui/icons-material/Edit'
import { projectsApi, Project } from '../api'

export default function ProjectsPage() {
  const [dialogOpen, setDialogOpen] = useState(false)
  const [editingProject, setEditingProject] = useState<Project | null>(null)
  const queryClient = useQueryClient()

  const { data: projects, isLoading, error } = useQuery({
    queryKey: ['projects'],
    queryFn: projectsApi.list,
  })

  const createMutation = useMutation({
    mutationFn: projectsApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['projects'] })
      setDialogOpen(false)
    },
  })

  const updateMutation = useMutation({
    mutationFn: ({ id, project }: { id: string; project: Partial<Project> }) =>
      projectsApi.update(id, project),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['projects'] })
      setDialogOpen(false)
      setEditingProject(null)
    },
  })

  const deleteMutation = useMutation({
    mutationFn: projectsApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['projects'] })
    },
  })

  const handleSubmit = (formData: FormData) => {
    const project: Omit<Project, 'id'> = {
      name: formData.get('name') as string,
      description: formData.get('description') as string,
    }

    if (editingProject?.id) {
      updateMutation.mutate({ id: editingProject.id, project })
    } else {
      createMutation.mutate(project)
    }
  }

  const handleEdit = (project: Project) => {
    setEditingProject(project)
    setDialogOpen(true)
  }

  const handleClose = () => {
    setDialogOpen(false)
    setEditingProject(null)
  }

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" mt={4}>
        <CircularProgress />
      </Box>
    )
  }

  if (error) {
    return (
      <Alert severity="error" sx={{ mt: 2 }}>
        Failed to load projects: {(error as Error).message}
      </Alert>
    )
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Projects
      </Typography>

      <Stack spacing={2}>
        {projects?.length === 0 && (
          <Typography color="text.secondary">
            No projects yet. Create your first project!
          </Typography>
        )}
        {projects?.map((project) => (
          <Card key={project.id}>
            <CardContent>
              <Box display="flex" justifyContent="space-between" alignItems="flex-start">
                <Box>
                  <Typography variant="h6">{project.name}</Typography>
                  {project.description && (
                    <Typography color="text.secondary" sx={{ mt: 1 }}>
                      {project.description}
                    </Typography>
                  )}
                </Box>
                <Box>
                  <IconButton onClick={() => handleEdit(project)} size="small">
                    <EditIcon />
                  </IconButton>
                  <IconButton
                    onClick={() => project.id && deleteMutation.mutate(project.id)}
                    size="small"
                    color="error"
                  >
                    <DeleteIcon />
                  </IconButton>
                </Box>
              </Box>
            </CardContent>
          </Card>
        ))}
      </Stack>

      <Fab
        color="primary"
        sx={{ position: 'fixed', bottom: 16, right: 16 }}
        onClick={() => setDialogOpen(true)}
      >
        <AddIcon />
      </Fab>

      <Dialog open={dialogOpen} onClose={handleClose} maxWidth="sm" fullWidth>
        <form
          onSubmit={(e) => {
            e.preventDefault()
            handleSubmit(new FormData(e.currentTarget))
          }}
        >
          <DialogTitle>{editingProject ? 'Edit Project' : 'New Project'}</DialogTitle>
          <DialogContent>
            <Stack spacing={2} sx={{ mt: 1 }}>
              <TextField
                name="name"
                label="Name"
                defaultValue={editingProject?.name || ''}
                required
                fullWidth
              />
              <TextField
                name="description"
                label="Description"
                defaultValue={editingProject?.description || ''}
                multiline
                rows={3}
                fullWidth
              />
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit" variant="contained">
              {editingProject ? 'Update' : 'Create'}
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </Box>
  )
}
