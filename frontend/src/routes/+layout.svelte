<script lang="ts">
	import '../app.css';
    import { Button } from "$lib/components/ui/button";
    import {
        Card,
        CardContent,
        CardDescription,
        CardFooter,
        CardHeader,
        CardTitle,
    } from "$lib/components/ui/card";
    import {
        Accordion,
        Root,
        AccordionContent,
        AccordionItem,
        AccordionTrigger,
    } from "$lib/components/ui/accordion";
    import { 
        ArrowRight, 
        Check, 
        ChevronDown, 
        Menu, 
        X 
    } from "lucide-svelte";
	
	let { children } = $props();
	let isMenuOpen = $state(false);
</script>

<!-- flex min-h-screen flex-col -->
<main class="container mx-auto px-4 py-8">
    
    <header class="container z-40 bg-background">
        <div class="flex h-20 items-center justify-between py-6">
            <div class="flex gap-6 md:gap-10">
                <a class="hidden items-center space-x-2 md:flex" href="/">
                    <span class="hidden font-bold sm:inline-block">GOUSS</span>
                </a>
                <nav class="hidden gap-6 md:flex">
                    {@render navLink("hidden gap-6 md:flex")}
                </nav>
            </div>
            <Button variant="outline" class="ml-auto mr-6" href="https://github.com/shadcn-ui/ui">
                GitHub
            </Button>
            <Button class="md:hidden" variant="ghost" size="icon" on:click={() => isMenuOpen = !isMenuOpen}>
                {#if isMenuOpen}
                    <X />
                {:else}
                    <Menu />
                {/if}
            </Button>
        </div>
    </header> 

    {#snippet navLink(cssClass)}
        <a class="flex items-center text-lg font-medium" href="#features">
            Features
        </a>
        <a class="flex items-center text-lg font-medium" href="#pricing">
            Pricing
        </a>
        <a class="flex items-center text-lg font-medium" href="#faq">
            FAQ
        </a>
    {/snippet}
    {#if isMenuOpen}
        <div class="container z-40 bg-background">
            <nav class="flex flex-col gap-4 p-4 md:hidden">
                {@render navLink()}
            </nav>
        </div>
    {/if}

    {@render children()}

    <footer class="border-t py-6 md:py-0">
        <div class="container flex flex-col items-center justify-between gap-4 md:h-24 md:flex-row">
            <p class="text-center text-sm leading-loose text-muted-foreground md:text-left">
                Built by <a
                    href="https://twitter.com/shadcn"
                    target="_blank"
                    rel="noreferrer"
                    class="font-medium underline underline-offset-4">shadcn</a
                >. Recreated in Svelte. The source code is available on
                <a
                    href="https://github.com/shadcn/ui"
                    target="_blank"
                    rel="noreferrer"
                    class="font-medium underline underline-offset-4">GitHub</a
                >.
            </p>
        </div>
    </footer>
</main>