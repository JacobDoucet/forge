import { useState } from 'react'
import { useQueryClient } from '@tanstack/react-query'
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
import { Task } from '../generated/model/task-model'
import {
  TaskStatus,
  TaskStatusPending,
  TaskStatusInProgress,
  TaskStatusCompleted,
} from '../generated/model/task-status-enum'
import {
  TaskPriority,
  TaskPriorityLow,
  TaskPriorityMedium,
  TaskPriorityHigh,
} from '../generated/model/task-priority-enum'
import {
  useSearchTasks,
  useCreateTask,
  useUpdateTask,
  useDeleteTask,
} from '../generated/react/tanstack-query/task-queries'

const statusColors: Record<TaskStatus, 'default' | 'primary' | 'success'> = {
  pending: 'default',
  in_progress: 'primary',
  completed: 'success',
  cancelled: 'default',
}

const statusLabels: Record<TaskStatus, string> = {
  pending: 'Pending',
  in_progress: 'In Progress',
  completed: 'Completed',
  cancelled: 'Cancelled',
}

const priorityColors: Record<TaskPriority, 'default' | 'warning' | 'error'> = {
  low: 'default',
  medium: 'warning',
  high: 'error',
  urgent: 'error',
}

export default function TasksPage() {
  const [dialogOpen, setDialogOpen] = useState(false)
  const [editingTask, setEditingTask] = useState<Task | null>(null)
  const queryClient = useQueryClient()

  const { data: tasksResponse, isLoading, error } = useSearchTasks({
    query: {},
  })

  const tasks = tasksResponse?.data || []

  const createMutation = useCreateTask({
    onAfterCommit: async () => {
      await queryClient.invalidateQueries({ queryKey: ['searchTasks'] })
      setDialogOpen(false)
    },
  })

  const updateMutation = useUpdateTask({
    onAfterCommit: async () => {
      await queryClient.invalidateQueries({ queryKey: ['searchTasks'] })
      setDialogOpen(false)
      setEditingTask(null)
    },
  })

  const deleteMutation = useDeleteTask({
    onAfterCommit: async () => {
      await queryClient.invalidateQueries({ queryKey: ['searchTasks'] })
    },
  })

  const handleSubmit = (formData: FormData) => {
    const statusValue = formData.get('status') as string
    const priorityValue = formData.get('priority') as string

    const task: Task = {
      id: editingTask?.id,
      title: formData.get('title') as string,
      description: formData.get('description') as string || undefined,
      status: (statusValue as TaskStatus) || TaskStatusPending,
      priority: priorityValue ? (priorityValue as TaskPriority) : undefined,
    }

    if (editingTask?.id) {
      updateMutation.mutate(task)
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
        Failed to load tasks: {error.message}
      </Alert>
    )
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Tasks
      </Typography>

      <Stack spacing={2}>
        {tasks.length === 0 && (
          <Typography color="text.secondary">
            No tasks yet. Create your first task!
          </Typography>
        )}
        {tasks.map(({ task }) => (
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
                    {task.status && (
                      <Chip
                        label={statusLabels[task.status]}
                        color={statusColors[task.status]}
                        size="small"
                      />
                    )}
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
                  defaultValue={editingTask?.status || TaskStatusPending}
                >
                  <MenuItem value={TaskStatusPending}>Pending</MenuItem>
                  <MenuItem value={TaskStatusInProgress}>In Progress</MenuItem>
                  <MenuItem value={TaskStatusCompleted}>Completed</MenuItem>
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
                  <MenuItem value={TaskPriorityLow}>Low</MenuItem>
                  <MenuItem value={TaskPriorityMedium}>Medium</MenuItem>
                  <MenuItem value={TaskPriorityHigh}>High</MenuItem>
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
