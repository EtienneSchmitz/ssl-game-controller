<script setup lang="ts">
import {computed, inject} from "vue";
import ToggleInput from "@/components/common/ToggleInput.vue";
import {useMatchStateStore} from "@/store/matchState";
import type {ControlApi} from "@/providers/controlApi/ControlApi";
import type {Team} from "@/proto/ssl_gc_common";

const props = defineProps<{
  team: Team,
}>()

const store = useMatchStateStore()
const control = inject<ControlApi>('control-api')

const model = computed(() => {
  return store.matchState.teamState![props.team].onPositiveHalf!
})

const updateValue = (value: boolean) => {
  control?.UpdateTeamState({
    forTeam: props.team,
    onPositiveHalf: value,
  })
}
</script>

<template>
  <ToggleInput
    :modelValue="model"
    label="Goal on positive field half"
    @update:model-value="updateValue"
  />
</template>
