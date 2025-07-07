<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from 'balzac/frontend/node_modules/@sveltejs/kit/src/runtime/app/navigation';

  onMount(() => {
    // After successful Google OAuth, the backend redirects to this page.
    // We can simply redirect to the home page or a dashboard.
    goto('/');
  });
</script>

<div class="container">
  <h1>Logging you in...</h1>
  <p>Please wait while we complete your login.</p>
</div>

<style>
  .container {
    max-width: 400px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    text-align: center;
  }

  h1 {
    color: #333;
    margin-bottom: 20px;
  }

  p {
    color: #666;
    margin-bottom: 20px;
  }
</style>