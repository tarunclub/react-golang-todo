import { useState, useEffect } from 'react';

function App() {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState('');

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    const response = await fetch('http://localhost:8000/tasks');
    const data = await response.json();
    setTasks(data);
  };

  const createTask = async () => {
    const response = await fetch('http://localhost:8000/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name: newTask, status: false }),
    });
    const data = await response.json();
    setTasks([...tasks, data]);
    setNewTask('');
  };

  const updateTask = async (id) => {
    const response = await fetch(`http://localhost:8000/tasks/${id}`, {
      method: 'PUT',
    });
    const data = await response.json();
    const updatedTasks = tasks.map((task) =>
      task.id === data.id ? data : task
    );
    setTasks(updatedTasks);
  };

  const deleteTask = async (id) => {
    await fetch(`http://localhost:8000/tasks/${id}`, {
      method: 'DELETE',
    });
    const updatedTasks = tasks.filter((task) => task.id !== id);
    setTasks(updatedTasks);
  };

  return (
    <div>
      <h1>TODO App</h1>
      <input
        type="text"
        value={newTask}
        onChange={(e) => setNewTask(e.target.value)}
      />
      <button onClick={createTask}>Add Task</button>
      <ul>
        {tasks.map((task) => (
          <li key={task.id}>
            {task.name} - {task.status ? 'Completed' : 'Pending'}
            <button onClick={() => updateTask(task.id)}>
              {task.status ? 'Mark as Pending' : 'Mark as Completed'}
            </button>
            <button onClick={() => deleteTask(task.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
