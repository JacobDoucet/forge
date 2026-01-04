import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import {
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Fab,
  FormControl,
  IconButton,
  InputLabel,
  MenuItem,
  Select,
  Stack,
  TextField,
  Typography,
  CircularProgress,
  Alert,
} from '@mui/material'
import AddIcon from '@mui/icons-material/Add'
import DeleteIcon from '@mui/icons-material/Delete'
import EditIcon from '@mui/icons-material/Edit'
import { tasksApi, Task } from '../api'

const statusColors: Record<string, 'default' | 'primary' | 'success'> = {
  todo: 'default',
  in_progress: 'primary',
  done: 'success',
}

const priorityColors: Record<string, 'default' | 'warning' | 'error'> = {
  low: 'default',
  medium: 'warning',
  high: 'error',
}

export default function TasksPage() {
  const [dialogOpen, setDialogOpen] = useState(false)
  const [editingTask, setEditingTask] = useState<Task | null>(null)
  const queryClient = useQueryClient()

  const { data: tasks, isLoading, error } = useQuery({
    queryKey: ['tasks'],
    queryFn: tasksApi.list,
  })

  const createMutation = useMutation({
    mutationFn: tasksApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] })
      setDialogOpen(false)
    },
  })

  const updateMutation = useMutation({
    mutationFn: ({ id, task }: { id: string; task: Partial<Task> }) =>
      tasksApi.update(id, task),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] })
      setDialogOpen(false)
      setEditingTask(null)
    },
  })

  const deleteMutation = useMutation({
    mutationFn: tasksApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['tasks'] })
    },
  })

  const handleSubmit = (formData: FormData) => {
    const task: Omit<Task, 'id'> = {
      title: formData.get('title') as string,
      description: formData.get('description') as string,
      status: (formData.get('status') as Task['status']) || 'todo',
      priority: formData.get('priority') as Task['priority'],
    }

    if (editingTask?.id) {
      updateMutation.mutate({ id: editingTask.id, task })
    } else {
      createMutation.mutate(task)
    }
  }

  const handleEdit = (task: Task) => {
    setEditingTask(task)
    setDialogOpen(true)
  }

  const handleClose = () => {
    setDialogOpen(false)
    setEditingTask(null)
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
        Failed to load tasks: {(error as Error).message}
      </Alert>
    )
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Tasks
      </Typography>

      <Stack spacing={2}>
        {tasks?.length === 0 && (
          <Typography color="text.secondary">
            No tasks yet. Create your first task!
          </Typography>
        )}
        {tasks?.map((task) => (
          <Card key={task.id}>
            <CardContent>
              <Box display="flex" justifyContent="space-between" alignItems="flex-start">
                <Box>
                  <Typography variant="h6">{task.title}</Typography>
                  {task.description && (
                    <Typography color="text.secondary" sx={{ mt: 1 }}>
                      {task.description}
                    </Typography>
                  )}
                  <Stack direction="row" spacing={1} sx={{ mt: 2 }}>
                    <Chip
                      label={task.status.replace('_', ' ')}
                      color={statusColors[task.status]}
                      size="small"
                    />
                    {task.priority && (
                      <Chip
                        label={task.priority}
                        color={priorityColors[task.priority]}
                        size="small"
                        variant="outlined"
                      />
                    )}
                  </Stack>
                </Box>
                <Box>
                  <IconButton onClick={() => handleEdit(task)} size="small">
                    <EditIcon />
                  </IconButton>
                  <IconButton
                    onClick={() => task.id && deleteMutation.mutate(task.id)}
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
          <DialogTitle>{editingTask ? 'Edit Task' : 'New Task'}</DialogTitle>
          <DialogContent>
            <Stack spacing={2} sx={{ mt: 1 }}>
              <TextField
                name="title"
                label="Title"
                defaultValue={editingTask?.title || ''}
                required
                fullWidth
              />
              <TextField
                name="description"
                label="Description"
                defaultValue={editingTask?.description || ''}
                multiline
                rows={3}
                fullWidth
              />
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  name="status"
                  label="Status"
                  defaultValue={editingTask?.status || 'todo'}
                >
                  <MenuItem value="todo">To Do</MenuItem>
                  <MenuItem value="in_progress">In Progress</MenuItem>
                  <MenuItem value="done">Done</MenuItem>
                </Select>
              </FormControl>
              <FormControl fullWidth>
                <InputLabel>Priority</InputLabel>
                <Select
                  name="priority"
                  label="Priority"
                  defaultValue={editingTask?.priority || ''}
                >
                  <MenuItem value="">None</MenuItem>
                  <MenuItem value="low">Low</MenuItem>
                  <MenuItem value="medium">Medium</MenuItem>
                  <MenuItem value="high">High</MenuItem>
                </Select>
              </FormControl>
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit" variant="contained">
              {editingTask ? 'Update' : 'Create'}
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </Box>
  )
}
