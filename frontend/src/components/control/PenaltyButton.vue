<script setup lang="ts">
import {computed, inject} from "vue";
import ControlButton from "@/components/control/ControlButton.vue";
import {useMatchStateStore} from "@/store/matchState";
import {Command_Type} from "@/proto/ssl_gc_state";
import type {ControlApi} from "@/providers/controlApi/ControlApi";
import type {Team} from "@/proto/ssl_gc_common";

const props = defineProps<{
  team: Team,
}>()

const store = useMatchStateStore()
const control = inject<ControlApi>('control-api')

const sendCommand = () => {
  control?.NewCommandForTeam(Command_Type.PENALTY, props.team)
}

const disable = computed(() => {
  return !store.isStop
})

</script>

<template>
  <ControlButton label="Penalty" :disable="disable" :action="sendCommand" :team="team"/>
</template>
