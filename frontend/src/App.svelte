<script>
  import { onMount } from 'svelte';
  import TodoList from './components/TodoList.svelte';
  import TodoForm from './components/TodoForm.svelte';

  let todos = [];
  let loading = true;

  const API_URL = '/api/todos';

  async function fetchTodos() {
    try {
      const response = await fetch(API_URL);
      todos = await response.json();
    } catch (error) {
      console.error('Error fetching todos:', error);
    } finally {
      loading = false;
    }
  }

  async function addTodo(title) {
    try {
      const response = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, completed: false })
      });
      const newTodo = await response.json();
      todos = [newTodo, ...todos];
    } catch (error) {
      console.error('Error adding todo:', error);
    }
  }

  async function toggleTodo(todo) {
    try {
      const response = await fetch(`${API_URL}/${todo.id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ ...todo, completed: !todo.completed })
      });
      const updatedTodo = await response.json();
      todos = todos.map(t => t.id === updatedTodo.id ? updatedTodo : t);
    } catch (error) {
      console.error('Error updating todo:', error);
    }
  }

  async function deleteTodo(id) {
    try {
      await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
      todos = todos.filter(t => t.id !== id);
    } catch (error) {
      console.error('Error deleting todo:', error);
    }
  }

  onMount(() => {
    fetchTodos();
  });
</script>

<main>
  <div class="container">
    <h1>Todo App</h1>
    <TodoForm on:add={(e) => addTodo(e.detail)} />
    {#if loading}
      <p>Loading...</p>
    {:else}
      <TodoList {todos} on:toggle={(e) => toggleTodo(e.detail)} on:delete={(e) => deleteTodo(e.detail)} />
    {/if}
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
  }

  main {
    padding: 2rem;
  }

  .container {
    max-width: 600px;
    margin: 0 auto;
    background: white;
    border-radius: 12px;
    box-shadow: 0 10px 40px rgba(0, 0, 0, 0.1);
    padding: 2rem;
  }

  h1 {
    margin: 0 0 2rem 0;
    color: #333;
    text-align: center;
    font-size: 2.5rem;
  }
</style>

