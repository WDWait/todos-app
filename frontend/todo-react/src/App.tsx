// src/App.tsx
import {useEffect, useState} from 'react'

interface Todo {
    id: number
    title: string
    completed: boolean
}

function App() {
    const [todos, setTodos] = useState<Todo[]>([])
    const [title, setTitle] = useState('')

    // 获取待办事项
    const fetchTodos = async () => {
        try {
            const res = await fetch('/api/todos')
            const data = await res.json()
            setTodos(data)
        } catch (error) {
            console.error('Failed to fetch todos:', error)
        }
    }

    // 创建待办
    const createTodo = async (e: React.FormEvent) => {
        e.preventDefault()
        if (!title.trim()) return

        try {
            const res = await fetch('/api/todos', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({title}),
            })
            const newTodo = await res.json()
            setTodos([...todos, newTodo])
            setTitle('')
        } catch (error) {
            console.error('Failed to create todo:', error)
        }
    }

    // 切换完成状态
    const toggleTodo = async (id: number, completed: boolean) => {
        try {
            // 根据 id 找到对应的 todo 项
            const todoToUpdate = todos.find(t => t.id === id);
            if (!todoToUpdate) {
                console.error('Todo not found');
                return;
            }

            await fetch(`/api/todos/${id}`, {
                method: 'PUT',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({
                    completed: !completed,
                    title: todoToUpdate.title
                }),
            })
            setTodos(todos.map(t => t.id === id ? {...t, completed: !completed} : t))
        } catch (error) {
            console.error('Failed to update todo:', error)
        }
    }

    // 删除待办
    const deleteTodo = async (id: number) => {
        try {
            await fetch(`/api/todos/${id}`, {method: 'DELETE'})
            setTodos(todos.filter(t => t.id !== id))
        } catch (error) {
            console.error('Failed to delete todo:', error)
        }
    }

    useEffect(() => {
        fetchTodos()
    }, [])

    return (
        <div className="min-h-screen bg-gray-50 py-8 px-4">
            <div className="max-w-md mx-auto bg-white rounded-lg shadow-md p-6">
                <h1 className="text-2xl font-bold text-center text-gray-800 mb-6">Todo App</h1>

                {/* 添加表单 */}
                <form onSubmit={createTodo} className="flex gap-2 mb-6">
                    <input
                        type="text"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        placeholder="Add a new task..."
                        className="flex-1 px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                    <button
                        type="submit"
                        className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition"
                    >
                        Add
                    </button>
                </form>

                {/* 列表 */}
                <ul className="space-y-3">
                    {todos.map(todo => (
                        <li
                            key={todo.id}
                            className="flex items-center justify-between p-3 border border-gray-200 rounded hover:bg-gray-50"
                        >
                            <div
                                onClick={() => toggleTodo(todo.id, todo.completed)}
                                className={`flex-1 cursor-pointer ${todo.completed ? 'line-through text-gray-500' : ''}`}
                            >
                                {todo.title}
                            </div>
                            <button
                                onClick={() => deleteTodo(todo.id)}
                                className="text-red-500 hover:text-red-700 text-sm ml-4"
                            >
                                Delete
                            </button>
                        </li>
                    ))}
                </ul>

                {todos.length === 0 && (
                    <p className="text-center text-gray-500 mt-4">No todos yet. Add one above!</p>
                )}
            </div>
        </div>
    )
}

export default App