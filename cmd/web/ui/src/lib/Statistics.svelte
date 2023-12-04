<script lang="ts">
  import { HardDrive, Indent, Rss } from 'lucide-svelte';
  import { onMount } from 'svelte';
  import { getHost } from './utils';
  import { indexersStore } from './store';

  let dbSize: string = '0 MiB';

  const fetchDBSize = async () => {
    const res = await fetch(getHost(`/management/db`));
    const data = await res.json();

    return `${(data / 1_000_000).toFixed(2)} MiB`;
  };

  onMount(async () => {
    dbSize = await fetchDBSize();
  });
</script>

<div class="stats shadow-lg 2xl:w-1/4 xl:w-1/3 lg:w-1/2 bg-base-200">
  <div class="stat">
    <div class="stat-figure text-primary">
      <HardDrive />
    </div>
    <div class="stat-title">Database size</div>
    <div class="stat-value text-primary">{dbSize}</div>
  </div>

  <div class="stat">
    <div class="stat-figure text-primary">
      <Indent />
    </div>
    <div class="stat-title">Indexers</div>
    {#await $indexersStore then indexers}
      <div class="stat-value text-primary">{indexers.length}</div>
    {/await}
  </div>
</div>
