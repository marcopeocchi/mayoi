<script lang="ts">
  import { HardDrive, Indent, Rss } from 'lucide-svelte';
  import { onMount } from 'svelte';

  let dbSize: string = '';
  let indexers: string[] = [];

  const fetchDBSize = async () => {
    const res = await fetch(`${window.location.href}management/db`);
    const data = await res.json();

    return `${(data / 1_000_000).toFixed(2)} MiB`;
  };

  const fetchIndexers = async () => {
    const res = await fetch(`${window.location.href}management/indexers`);
    const data = await res.json();

    return data;
  };

  onMount(async () => {
    const [_dbSize, _indexers] = await Promise.all([
      fetchDBSize(),
      fetchIndexers(),
    ]);
    dbSize = _dbSize;
    indexers = _indexers;
  });
</script>

<div class="bg-neutral-100 border w-screen mx-8 p-8 text-sm">
  <div class="flex flex-row gap-2 items-center">
    <Rss />
    <p>RSS/Torznab address:</p>
    <button
      class="bg-neutral-200 px-1"
      on:click={() => window.open(`${window.location.href}api`)}
    >
      {window.location.href}api
    </button>
  </div>

  <div class="flex flex-row gap-2 items-center mt-4">
    <HardDrive />
    <p>Database size:</p>
    <p class="bg-neutral-200 px-1">
      {dbSize}
    </p>
  </div>

  <div class="flex flex-row gap-2 mt-4">
    <Indent />
    <p>Indexers:</p>
    <ul>
      {#each indexers as index}
        <li class="bg-neutral-200 px-1">
          &bull; {index}
        </li>
      {/each}
    </ul>
  </div>
</div>
