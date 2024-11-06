<script lang="ts">
    import { ArrowRight, Check, ChevronDown, Menu, X, Link, Zap, Shield } from "lucide-svelte";
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
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
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
    let url = "";
    let customUrl = "";
    let shortenedUrl = "";

    function handleSubmit() {
        // Here you would typically make an API call to your backend to shorten the URL
        // For this example, we'll just simulate it
        shortenedUrl = `https://short.url/${Math.random().toString(36).substr(2, 5)}`;
    }
</script>

<section class="space-y-6 pb-8 pt-6 md:pb-12 md:pt-10 lg:py-32">
    <div class="text-center mb-12">
        <h1 class="text-4xl font-bold mb-4">Gouss</h1>
        <p class="text-xl text-muted-foreground">
            Transform long, complex URLs into short, shareable links with
            just a click.
        </p>
        <p class="text-xl text-muted-foreground">
            Boost your online presence and track your link
            performance effortlessly.
        </p>
    </div>
</section>

<section class="grid md:grid-cols-2 gap-8 items-center mb-16">
    <div class="container flex flex-col gap-6 py-8 md:max-w-[64rem] md:py-12 lg:py-24">
        <div class="mx-auto flex w-full flex-col gap-4 md:max-w-[58rem]">
            <h2 class="font-heading text-3xl leading-[1.1] sm:text-3xl md:text-6xl">
                Get Started Today
            </h2>
            <p class="max-w-[85%] leading-normal text-muted-foreground sm:text-lg sm:leading-7">
                Shorten your links, expand your reach
            </p>
        </div>
        <div class="flex flex-col gap-4 sm:flex-row">
            <Button size="lg">
                Get Started
                <ArrowRight class="ml-2" />
            </Button>
        </div>
    </div>
    <Card>
        <CardHeader>
            <CardTitle>Shorten Your URL</CardTitle>
            <CardDescription
                >Paste your long URL and we'll do the magic</CardDescription
            >
        </CardHeader>
        <CardContent>
            <form on:submit|preventDefault={handleSubmit} class="space-y-4">
                <Input
                    type="url"
                    placeholder="Enter your long URL here"
                    bind:value={url}
                    required
                />
                <Input
                    type="input"
                    placeholder="Custom Alias (Optional)"
                    bind:value={customUrl}
                />
                <Button type="submit" class="w-full">Shorten URL</Button>
            </form>
        </CardContent>
        {#if shortenedUrl}
            <CardFooter>
                <p class="text-sm text-muted-foreground">
                    Your shortened URL:
                </p>
                <a
                    href={shortenedUrl}
                    class="text-primary font-medium"
                    target="_blank"
                    rel="noopener noreferrer">{shortenedUrl}</a
                >
            </CardFooter>
        {/if}
    </Card>
</section>

<section class="space-y-6 pb-8 pt-6 md:pb-12 md:pt-10 lg:py-32">
    <div class="mb-16">
        <h2 class="text-3xl font-semibold text-center mb-8">
            Why Choose QuickLink?
        </h2>
        <div class="grid md:grid-cols-3 gap-8">
            <Card>
                <CardHeader>
                    <Zap class="mb-2 text-primary" />
                    <CardTitle>Lightning Fast</CardTitle>
                </CardHeader>
                <CardContent>
                    <p>
                        Generate short links in seconds and share them instantly
                        across all platforms.
                    </p>
                </CardContent>
            </Card>
            <Card>
                <CardHeader>
                    <Link class="mb-2 text-primary" />
                    <CardTitle>Custom Links</CardTitle>
                </CardHeader>
                <CardContent>
                    <p>
                        Create memorable, branded short links that reflect your
                        business or campaign.
                    </p>
                </CardContent>
            </Card>
            <Card>
                <CardHeader>
                    <Shield class="mb-2 text-primary" />
                    <CardTitle>Secure & Reliable</CardTitle>
                </CardHeader>
                <CardContent>
                    <p>
                        Rest easy knowing your links are protected and always
                        accessible.
                    </p>
                </CardContent>
            </Card>
        </div>
    </div>
</section>

<section class="container flex flex-col gap-6 py-8 md:max-w-[64rem] md:py-12 lg:py-24">
    <div class="mx-auto flex w-full flex-col gap-4 md:max-w-[58rem]">
        <h2 class="font-heading text-3xl leading-[1.1] sm:text-3xl md:text-6xl">
            Testimonials
        </h2>
        <p class="max-w-[85%] leading-normal text-muted-foreground sm:text-lg sm:leading-7">
            See what others are saying about this template.
        </p>
    </div>
    <div class="grid w-full grid-cols-1 gap-4 md:grid-cols-2">
        <Card>
            <CardHeader>
                <CardTitle>Great template!</CardTitle>
            </CardHeader>
            <CardContent>
                <p>
                    "This template has everything I need to build my Svelte
                    app. Highly recommended!"
                </p>
            </CardContent>
            <CardFooter>
                <p class="text-sm text-muted-foreground">- Jane Doe</p>
            </CardFooter>
        </Card>
        <Card>
            <CardHeader>
                <CardTitle>Saved me so much time</CardTitle>
            </CardHeader>
            <CardContent>
                <p>
                    "I was able to build my app in record time thanks to
                    this template. It's a game-changer!"
                </p>
            </CardContent>
            <CardFooter>
                <p class="text-sm text-muted-foreground">- John Smith</p>
            </CardFooter>
        </Card>
    </div>
</section>

<section id="pricing" class="container flex flex-col gap-6 py-8 md:max-w-[64rem] md:py-12 lg:py-24">
    <div class="mx-auto flex w-full flex-col gap-4 md:max-w-[58rem]">
        <h2 class="font-heading text-3xl leading-[1.1] sm:text-3xl md:text-6xl">
            Simple, transparent pricing
        </h2>
        <p class="max-w-[85%] leading-normal text-muted-foreground sm:text-lg sm:leading-7">
            Choose the plan that's right for you
        </p>
    </div>
    <div class="grid w-full grid-cols-1 gap-4 md:grid-cols-2">
        {#each data.pricingPlans as plan}
            <Card>
                <CardHeader>
                    <CardTitle>{plan.name}</CardTitle>
                    <CardDescription>{plan.description}</CardDescription>
                </CardHeader>
                <CardContent>
                    <p class="text-4xl font-bold">
                        {plan.price}<span class="text-lg font-normal"
                            >/month</span
                        >
                    </p>
                    <ul class="mt-4 space-y-2">
                        {#each plan.features as feature}
                            <li class="flex items-center">
                                <Check class="mr-2 h-4 w-4 text-primary" />
                                {feature}
                            </li>
                        {/each}
                    </ul>
                </CardContent>
                <CardFooter>
                    <Button class="w-full">Get Started</Button>
                </CardFooter>
            </Card>
        {/each}
    </div>
</section>

<section id="faq" class="container flex flex-col gap-6 py-8 md:max-w-[64rem] md:py-12 lg:py-24">
    <div class="mx-auto flex w-full flex-col gap-4 md:max-w-[58rem]">
        <h2 class="font-heading text-3xl leading-[1.1] sm:text-3xl md:text-6xl">
            Frequently Asked Questions
        </h2>
        <p class="max-w-[85%] leading-normal text-muted-foreground sm:text-lg sm:leading-7">
            Find answers to common questions about our template.
        </p>
    </div>
    <Root class="w-full">
        {#each data.faqs as faq, index}
            <AccordionItem value={`item-${index}`}>
                <AccordionTrigger>{faq.question}</AccordionTrigger>
                <AccordionContent>{faq.answer}</AccordionContent>
            </AccordionItem>
        {/each}
    </Root>
</section>

<style>
    :global(body) {
        @apply bg-background text-foreground;
    }
</style>
