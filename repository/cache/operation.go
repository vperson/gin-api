package cache

import "time"

func (c *Client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.redisClient.Set(c.ctx, key, value, expiration).Err()
}

func (c *Client) Get(key string) (string, error) {
	return c.redisClient.Get(c.ctx, key).Result()
}

func (c *Client) Del(key string) error {
	return c.redisClient.Del(c.ctx, key).Err()
}

// Like searches for keys matching a pattern and returns their values
func (c *Client) Like(pattern string) (map[string]string, error) {
	var cursor uint64
	results := make(map[string]string)
	for {
		// Use SCAN command to find matching keys
		keys, nextCursor, err := c.redisClient.Scan(c.ctx, cursor, pattern, 0).Result()
		if err != nil {
			return nil, err
		}

		// Fetch values for the found keys
		for _, key := range keys {
			value, err := c.Get(key)
			if err != nil {
				return nil, err
			}
			results[key] = value
		}

		// Exit loop if the cursor is 0 (end of scan)
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return results, nil
}
