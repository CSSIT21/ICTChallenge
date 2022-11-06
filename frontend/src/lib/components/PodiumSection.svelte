<script lang="ts">
        import PodiumComponent from "./PodiumComponent.svelte";
</script>

<div class="relative h-[653px] flex justify-center gap-10 items-end">
        <div class="absolute bottom-0 w-screen scale-90">
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2  h-[256px] w-[512px] rounded-t-[384px]" />
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2 h-[384px] w-[768px] rounded-t-[576px]" />
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2 h-[512px] w-[1024px] rounded-t-[768px]" />
                <div class="bg-gradient-to-b from-[rgb(255,255,255,0.08)] to-[rgb(255,255,255,0)] absolute bottom-0 left-1/2 -translate-x-1/2 h-[640px] w-[1280px] rounded-t-[960px]" />
        </div>
        <PodiumComponent team="cympati" order={2} score={1000} />
        <PodiumComponent team="cympati" order={1} score={20000} />
        <PodiumComponent team="cympati" order={3} score={500} />
        
</div>
