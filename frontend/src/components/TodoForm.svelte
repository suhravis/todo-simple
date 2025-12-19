<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  let title = '';

  function handleSubmit() {
    if (title.trim()) {
      dispatch('add', title.trim());
      title = '';
    }
  }
</script>

<form on:submit|preventDefault={handleSubmit}>
  <div class="form-group">
    <input
      type="text"
      bind:value={title}
      placeholder="Add a new todo..."
      class="input"
    />
    <button type="submit" class="button">Add</button>
  </div>
</form>

<style>
  .form-group {
    display: flex;
    gap: 0.5rem;
    margin-bottom: 1.5rem;
  }

  .input {
    flex: 1;
    padding: 0.75rem;
    border: 2px solid #e0e0e0;
    border-radius: 8px;
    font-size: 1rem;
    transition: border-color 0.2s;
  }

  .input:focus {
    outline: none;
    border-color: #667eea;
  }

  .button {
    padding: 0.75rem 1.5rem;
    background: #667eea;
    color: white;
    border: none;
    border-radius: 8px;
    font-size: 1rem;
    cursor: pointer;
    transition: background 0.2s;
  }

  .button:hover {
    background: #5568d3;
  }

  .button:active {
    transform: scale(0.98);
  }
</style>

