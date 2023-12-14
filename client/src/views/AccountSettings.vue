<template>
    <section>
<h3>Account Settings</h3>
<RadioButtons :radios="radios" />
    </section>
</template>
<script setup lang="ts">
import RadioButtons from '@/components/RadioButtons.vue';

const radios = [
    {value: "man"},
    {value: "woman"},
    {value: "other"},
]
</script>