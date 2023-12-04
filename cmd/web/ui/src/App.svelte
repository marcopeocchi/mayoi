<script lang="ts">
  import Indexers from './lib/Indexers.svelte';
  import Navbar from './lib/Navbar.svelte';
  import Statistics from './lib/Statistics.svelte';
</script>

<main class="min-h-screen">
  <Navbar />
  <div class="p-8 flex flex-col gap-8">
    <Statistics />
    <Indexers />
  </div>
</main>
