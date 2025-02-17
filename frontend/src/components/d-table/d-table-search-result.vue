<template>
  <div
    v-bind="$attrs"
    v-if="fetching && data && data[objectName]?.edges?.length > 0"
    v-for="i in variables.limit"
    :key="i"
    class="animate-pulse"
  >
    <div v-for="column in columns" :key="column.key" class="border-b border-neutral-100 px-8 py-3">
      <div class="h-3 rounded-full bg-neutral-100"></div>
    </div>
  </div>
  <div
    v-bind="$attrs"
    v-if="!fetching && data && data[objectName]?.edges?.length > 0"
    v-for="row in data[objectName]?.edges"
    :key="row.id"
    ref="items"
    :draggable="draggable"
    class="group/row grid border-b transition-colors hover:bg-neutral-50"
    :class="
      dragoverItem === row.id && draggingItem !== row.id
        ? 'border-blue-500 bg-blue-100'
        : 'border-b-neutral-100 border-transparent'
    "
    @dragstart="(event) => dragStart(event, row)"
    @dragend="dragend"
    @drop="(event) => drop(event, row)"
    @dragover.prevent="(event) => dragover(event, row)"
  >
    <slot :row="row" :fetching="fetching"></slot>
  </div>
  <slot
    name="empty"
    v-if="
      (!fetching && !data && pageVariables.offset === 0) ||
      (!fetching && data && data[objectName]?.edges?.length === 0 && pageVariables.offset === 0)
    "
  />
</template>

<script lang="ts" setup>
import { ref, watch, toRef } from "vue"
import { useMoveFileMutation } from "@/gql/mutations/files/moveFile"
import type { File } from "@/gql/schema"
import { useQuery } from "@urql/vue"

const props = defineProps([
  "query",
  "variables",
  "objectName",
  "columns",
  "additionalTypenames",
  "draggable",
  "dragDataType"
])
const pageVariables = toRef(props, "variables")

const items = ref<HTMLElement[]>([])
const dragoverItem = ref<string | null>(null)
const draggingItem = ref<string | null>(null)
const draggable = toRef(props, "draggable")
const dragDataType = toRef(props, "dragDataType")

function dragStart(event: DragEvent, row: any) {
  if (!draggable.value) return

  event.dataTransfer?.setData("dokedu/vnd.dokedu-drive-file", row.id)
  draggingItem.value = row.id
}

async function drop(event: DragEvent, row: File) {
  if (!draggable.value) return

  const data: string | undefined = event.dataTransfer?.getData(dragDataType.value)
  if (!data) return

  if (!row) return
  if (!row.id) return

  if (row.fileType !== "folder") return
  if (row.id === data) return

  await moveFile({
    input: {
      id: data,
      targetId: row.id
    }
  })
}

function dragover(_: DragEvent, row: File) {
  if (!draggable.value) return

  dragoverItem.value = row.id
}

function dragend() {
  if (!draggable.value) return

  dragoverItem.value = null
  draggingItem.value = null
}

const { data, fetching } = useQuery({
  query: props.query,
  variables: pageVariables,
  context: {
    additionalTypenames: (props.additionalTypenames || []) as string[]
  }
})

const { executeMutation: moveFile } = useMoveFileMutation()

watch(data, () => {
  if (!data.value) return
  pageVariables.value.nextPage = data.value[props.objectName]?.pageInfo.hasNextPage
})
</script>
