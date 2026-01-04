const API_BASE = '/api'

export interface Task {
  id?: string
  title: string
  description?: string
  status: 'todo' | 'in_progress' | 'done'
  priority?: 'low' | 'medium' | 'high'
  dueDate?: string
  tags?: string[]
  assigneeId?: string
}

export interface Project {
  id?: string
  name: string
  description?: string
}

async function fetchJson<T>(url: string, options?: RequestInit): Promise<T> {
  const response = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
  
  if (!response.ok) {
    const error = await response.text()
    throw new Error(error || `HTTP error! status: ${response.status}`)
  }
  
  if (response.status === 204) {
    return {} as T
  }
  
  return response.json()
}

// Tasks API
export const tasksApi = {
  list: () => fetchJson<Task[]>(`${API_BASE}/tasks/`),
  
  get: (id: string) => fetchJson<Task>(`${API_BASE}/tasks/${id}`),
  
  create: (task: Omit<Task, 'id'>) =>
    fetchJson<Task>(`${API_BASE}/tasks/`, {
      method: 'POST',
      body: JSON.stringify(task),
    }),
  
  update: (id: string, task: Partial<Task>) =>
    fetchJson<Task>(`${API_BASE}/tasks/${id}`, {
      method: 'PUT',
      body: JSON.stringify(task),
    }),
  
  delete: (id: string) =>
    fetchJson<void>(`${API_BASE}/tasks/${id}`, {
      method: 'DELETE',
    }),
}

// Projects API
export const projectsApi = {
  list: () => fetchJson<Project[]>(`${API_BASE}/projects/`),
  
  get: (id: string) => fetchJson<Project>(`${API_BASE}/projects/${id}`),
  
  create: (project: Omit<Project, 'id'>) =>
    fetchJson<Project>(`${API_BASE}/projects/`, {
      method: 'POST',
      body: JSON.stringify(project),
    }),
  
  update: (id: string, project: Partial<Project>) =>
    fetchJson<Project>(`${API_BASE}/projects/${id}`, {
      method: 'PUT',
      body: JSON.stringify(project),
    }),
  
  delete: (id: string) =>
    fetchJson<void>(`${API_BASE}/projects/${id}`, {
      method: 'DELETE',
    }),
}
