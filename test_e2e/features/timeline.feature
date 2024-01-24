Feature: Timeline
  In order to see the recent tweets of my followed users
  As an user
  I want to see the timeline

  Scenario: See the timeline
    Given I follow users
      | 1 |
      | 2 |
      | 3 |
# And user 1 tweets "Hello i'm user 1"
# And user 2 tweets "Hello i'm user 2"
# And user 3 tweets "Hello i'm user 3"
# And user 4 tweets "Hello i'm user 4"
# When I see the timeline
# Then I see tweets
#   | Hello i'm user 1 |
#   | Hello i'm user 2 |
#   | Hello i'm user 3 |