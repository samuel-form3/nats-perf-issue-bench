# nats-perf-issue-bench

## How-to Reproduce

### With NATS 2.7.1

```bash
make setup-environment-2.7.1
```

```bash
make run-performance-test
```

### With NATS 2.7.4

```bash
make setup-environment-2.7.4
```

```bash
make run-performance-test
```

## What happens

When running NATS 2.7.4, we start observing very high latencies when publishing messages.
