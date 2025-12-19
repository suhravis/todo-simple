<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  export let todos = [];
</script>

<div class="todo-list">
  {#if todos.length === 0}
    <p class="empty">No todos yet. Add one above!</p>
  {:else}
    {#each todos as todo (todo.id)}
      <div class="todo-item" class:completed={todo.completed}>
        <input
          type="checkbox"
          checked={todo.completed}
          on:change={() => dispatch('toggle', todo)}
          class="checkbox"
        />
        <span class="title">{todo.title}</span>
        <button
          on:click={() => dispatch('delete', todo.id)}
          class="delete-button"
        >
          ×
        </button>
      </div>
    {/each}
  {/if}
</div>

<style>
  .todo-list {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .todo-item {
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 1rem;
    background: #f8f9fa;
    border-radius: 8px;
    transition: all 0.2s;
  }

  .todo-item:hover {
    background: #e9ecef;
  }

  .todo-item.completed {
    opacity: 0.6;
  }

  .todo-item.completed .title {
    text-decoration: line-through;
    color: #6c757d;
  }

  .checkbox {
    width: 20px;
    height: 20px;
    cursor: pointer;
  }

  .title {
    flex: 1;
    font-size: 1rem;
    color: #333;
  }

  .delete-button {
    width: 32px;
    height: 32px;
    border: none;
    background: #dc3545;
    color: white;
    border-radius: 50%;
    font-size: 1.5rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background 0.2s;
    line-height: 1;
  }

  .delete-button:hover {
    background: #c82333;
  }

  .empty {
    text-align: center;
    color: #6c757d;
    padding: 2rem;
    font-style: italic;
  }
</style>

