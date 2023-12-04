<script lang="ts">
  import type { Indexer } from '../types';
  import { indexersStore } from './store';
  import { getHost } from './utils';
</script>

<div class="overflow-x-auto bg-base-200 rounded-lg shadow-lg">
  <table class="table">
    <thead>
      <tr>
        <th></th>
        <th>Indexer</th>
        <th>Torznab/RSS URL</th>
      </tr>
    </thead>
    <tbody>
      {#await $indexersStore}
        <span class="loading loading-ring loading-xs" />
      {:then indexers}
        {#each indexers as indexer, idx}
          <tr class="bg-base-200">
            <th>{idx + 1}</th>
            <td>{indexer.url}</td>
            <td>
              <button
                class="btn btn-neutral"
                on:click={() => window.open(getHost(indexer.path))}
              >
                {getHost(indexer.path)}
              </button>
            </td>
          </tr>
        {/each}
      {/await}
    </tbody>
  </table>
</div>
