import { readable } from 'svelte/store'
import { getHost } from './utils'
import type Indexers from './Indexers.svelte'

const fetchIndexers = async () => {
  const res = await fetch(getHost('/management/indexers'))
  const data = await res.json()

  return data
}

export const indexersStore = readable<Promise<Indexers[]>>(fetchIndexers())