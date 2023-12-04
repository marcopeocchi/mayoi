<script lang="ts">
  import Navbar from './lib/Navbar.svelte';
  import Statistics from './lib/Statistics.svelte';
</script>

<main class="min-h-screen dark:bg-neutral-900 dark:text-neutral-100">
  <Navbar />
  <div class="flex items-center justify-center mt-8">
    <Statistics />
  </div>
</main>
